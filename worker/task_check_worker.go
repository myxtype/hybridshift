package worker

import (
	"context"
	"errors"
	"frame/model"
	"frame/pkg/logger"
	"frame/pkg/twt"
	"frame/store/db"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/shopspring/decimal"
	"time"
)

type TaskCheckWorker struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
}

var taskCheckSkipErr = errors.New("skip")

func NewTaskCheckWorker() *TaskCheckWorker {
	ctx, cancel := context.WithCancel(context.Background())
	return &TaskCheckWorker{
		ctx:       ctx,
		ctxCancel: cancel,
	}
}

func (w *TaskCheckWorker) Start() {
	for {
		select {
		case <-time.After(15 * time.Minute):
			var (
				id    = uint(0)
				limit = 50
			)

			for {
				subtasks, err := db.Shared().GetValidTaskSubtasksLimit(id, limit)
				if err != nil {
					logger.Sugar.Error(err)
					break
				}
				if len(subtasks) == 0 {
					break
				}

				for _, n := range subtasks {
					id = n.ID

					if err := w.CheckSubtask(n); err != nil {
						if err == taskCheckSkipErr {
							continue
						}
						logger.Sugar.Error(err)
					}

					time.Sleep(600 * time.Millisecond)
				}

				if len(subtasks) < limit {
					break
				}
			}

		case <-w.ctx.Done():
			return
		}
	}
}

func (w *TaskCheckWorker) Stop() {
	w.ctxCancel()
}

func (w *TaskCheckWorker) CheckSubtask(subtask *model.TaskSubtask) error {
	// 15分钟之类检查过，就跳过
	if subtask.CheckedAt.Before(time.Now().Add(-15 * time.Minute)) {
		return taskCheckSkipErr
	}

	// 事务
	tx, err := db.Shared().BeginTx()
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	// 获取此用户所有的子任务
	subtasks, err := tx.GetValidTaskSubtasksByUser(subtask.UserId)
	if err != nil {
		return err
	}
	if len(subtasks) == 0 {
		return nil
	}

	oauth, err := tx.GetUserOauth(subtask.UserId)
	if err != nil {
		return err
	}
	if oauth == nil {
		logger.Sugar.Warnf("没有找到用户的推特认证：%v", subtask.UserId)
		return nil
	}

	client := twt.ClientV2(oauth.OauthToken, oauth.OauthSecret)

	// 抓取用户最近100条Timeline
	rsp, err := client.UserTweetTimeline(w.ctx, oauth.OauthId, twitter.UserTweetTimelineOpts{
		MaxResults: 100,
		TweetFields: []twitter.TweetField{
			twitter.TweetFieldReferencedTweets,
			twitter.TweetFieldPublicMetrics,
			twitter.TweetFieldEntities,
		},
		Expansions: []twitter.Expansion{
			twitter.ExpansionReferencedTweetsID,
			twitter.ExpansionEntitiesMentionsUserName,
		},
	})
	if err != nil {
		return err
	}

	tweets := rsp.Raw.Tweets

	if len(tweets) == 0 {
		return nil
	}

	var taskIds []uint
	for _, n := range subtasks {
		taskIds = append(taskIds, n.TaskId)
	}

	tasks, err := tx.GetTasksByIdForUpdate(taskIds)
	if err != nil {
		return err
	}

	var (
		validTaskIDs []uint
		now          = time.Now()
	)

	for _, n := range subtasks {
		task := tasks[n.TaskId]

		// 判断任务状态
		if task == nil || task.Status != model.TaskStatusOn {
			n.Valid = false
			if err := tx.UpdateTaskSubtaskValid(n); err != nil {
				return err
			}
			continue
		}

		validTaskIDs = append(validTaskIDs, n.TaskId)

		for _, tweet := range tweets {
			isAt := false
			for _, em := range tweet.Entities.Mentions {
				if em.UserName == task.TwitterUserName {
					isAt = true
					break
				}
			}
			if !isAt {
				continue
			}

			innerTweet, err := tx.GetTaskTweetByTweetID(tweet.ID)
			if err != nil {
				return err
			}

			// 判断推特类型
			tweetType := task.CheckTweetType(tweet.ReferencedTweets)

			if innerTweet == nil {
				innerTweet = &model.TaskTweet{
					UserId:     n.UserId,
					TaskId:     n.TaskId,
					SubtaskId:  n.ID,
					TweetID:    tweet.ID,
					Type:       tweetType,
					TweetURL:   "",
					Views:      0,
					Likes:      0,
					Replies:    0,
					Quotes:     0,
					Retweets:   0,
					TotalPoint: decimal.Zero,
				}
				if err := tx.AddTaskTweet(innerTweet); err != nil {
					return err
				}
			}
		}

		n.CheckedAt = now
		if err := tx.UpdateTaskSubtaskCheck(n); err != nil {
			return err
		}
	}

	// 获取任务推特
	taskTweets, err := tx.GetTaskTweetsByUser(subtask.UserId, validTaskIDs)
	if err != nil {
		return err
	}

	if len(taskTweets) > 0 {
		var tweetIds []string
		for _, n := range taskTweets {
			if !n.Deleted {
				tweetIds = append(tweetIds, n.TweetID)
			}
		}

		lookupRsp, err := client.TweetLookup(w.ctx, tweetIds, twitter.TweetLookupOpts{})
		if err != nil {
			return err
		}

		tweetDic := lookupRsp.Raw.TweetDictionaries()

		for _, n := range taskTweets {
			if n.Deleted {
				continue
			}

			tweet, found := tweetDic[n.TweetID]
			if !found {
				n.Deleted = true
				if err := tx.UpdateTaskTweetDelete(n); err != nil {
					return err
				}
			}

			addPoint := n.SetUpdateMetrics(tweet.Tweet.PublicMetrics)
			if addPoint.IsPositive() {
				// todo
			}
		}
	}

	// 提交
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
