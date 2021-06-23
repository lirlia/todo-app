package main

import (
	"net/http"

	"todo.app/controller"
	"todo.app/model"

	"github.com/gin-gonic/gin"
)

/*
// MapからStructへの変換
func MapToStruct(m map[string]string) logMessage {
	reqtime, err := strconv.ParseFloat(m["reqtime"], 64)
	if err != nil {
		panic(err)
	}

	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, m["time"])

	if err != nil {
		panic(err)
	}

	return logMessage{m["id"], t, m["level"], m["method"], m["uri"], reqtime, time.Now()}
}


// ログをDBにBatchINSERTする
// DB_INSERT_BATCHSIZEで同時に格納するサイズを調整
func BatchInsertToDB(db *gorm.DB, l []logMessage) {
	db.CreateInBatches(l, DB_INSERT_BATCHSIZE)
}

func main() {

	// dbへの接続
	db := connectDB()

	// log用のテーブルの作成を行う
	createTable(db)

	// LOG_DIRディレクトリに存在する圧縮ファイルを取得する
	filelist, err := ioutil.ReadDir(LOG_DIR)
	fileCount := len(filelist)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	// ファイルを解凍してDBに格納する関数
	mainFunc := func(l int) {
		f := filelist[l]

		// wgをデクリメントしgo routineの完了を示す
		defer wg.Done()

		// 対象のファイルを開く
		filename := LOG_DIR + "/" + f.Name()
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// gzを解凍&展開する
		gzipReader, err := gzip.NewReader(file)
		if err != nil {
			panic(err)
		}
		defer gzipReader.Close()

		// 解凍したデータを読み込み、logMessage構造体に変換する
		scanner := bufio.NewScanner(gzipReader)
		var logs = []logMessage{}

		//　読み込んだデータを1行づつ確認
		for scanner.Scan() {

			// ログをLTSV→logMessage構造体に変換する
			// 一度LTSV→Mapに変換
			record, err := ltsv.ParseLineAsMap([]byte(scanner.Text()), nil)
			if err != nil {
				panic(err)
			}

			// []logMessage{} 型につっこみまとめてINSERTできるようにしておく
			logs = append(logs, MapToStruct(record))
		}
		// メモリに読み込んだログをINSERTする
		BatchInsertToDB(db, logs)

		fmt.Println(filename + " insert succeed.")
	}

	// go routineでファイルごとに並行処理
	for l := 0; l < fileCount; l++ {
		l := l
		wg.Add(1)
		go mainFunc(l)
	}

	// go routineの完了まで待つ
	wg.Wait()
	fmt.Println("finish")
}

func setTasktoList(task_id int, name string, msg string, done bool, user_id int) (r *Task) {
	r = new(model.Task)
	r.TaskID = task_id
	r.Name = name
	r.Message = msg
	r.Done = done
	r.UserID = user_id
	return r
}
*/

//type TaskList = []*Task

func setupRouter() *gin.Engine {

	//	var tasks TaskList
	//var task = model.Task{}

	//	tasks = append(tasks, setTasktoList(1, "これはテストです1", "ああああああ", false, 1))
	//	tasks = append(tasks, setTasktoList(2, "これはテストです2", "ああああああ", true, 1))
	//	tasks = append(tasks, setTasktoList(3, "これはテストです3", "ああああああ", true, 1))

	r := gin.Default()
	r.Static("/static", "../static")
	r.LoadHTMLGlob("../templates/*.*")

	task_r := r.Group("/task")
	{
		v1 := task_r.Group("/v1")
		{
			v1.POST("/add", controller.TaskAdd)
			v1.GET("/list", controller.TaskList)
			v1.PUT("/update", controller.TaskUpdate)
			v1.DELETE("/delete", controller.TaskDelete)
		}
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/mypage", func(c *gin.Context) {
		var user = model.User{
			Name:     c.PostForm("username"),
			Password: c.PostForm("password"),
		}

		c.HTML(http.StatusOK, "mypage.html", gin.H{
			"username": user.Name,
			//"tasks":    db.Find(&task),
		})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
