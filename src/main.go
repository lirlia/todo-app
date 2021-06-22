package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
const (
	DB_NAME                = "todo"
	DB_USER                = "root"
	DB_PASS                = "password"
	DB_HOST                = "db"
	DB_PORT                = "3306"
	DB_ACCOUNT_TABLE_NAME  = "account"
	DB_TODOLIST_TABLE_NAME = "todolist"
	DB_INSERT_BATCHSIZE    = 3000

	// SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定します
	DB_MAX_IDLE_CONN = 10

	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定します
	DB_MAX_OPEN_CONN = 100

	// SetConnMaxLifetimeは再利用され得る最長時間を設定します
	DB_MAX_LIFETIME = time.Hour
)

type Account struct {
	AccountID int `gorm:"primaryKey"`
	Name      string
	Password  string
	CreatedAt time.Time
}

type Task struct {
	TaskID    string `gorm:"primaryKey"`
	Name      string
	Done      bool
	Message   string
	AccountID string `gorm:"foreignKey"`
	CreatedAt time.Time
}

// テーブル名を変換するために定義
type Tabler interface {
	TableName() string
}

// テーブル名を決定する
func (logMessage) TableName() string {
	return DB_TABLE_NAME
}

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

// DBへの接続を行う関数
func connectDB() *gorm.DB {
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// トランザクションの中で整合性を担保するための設定
		// 今回は特定のレコードに対して複数の操作を行うことはないため性能向上を鑑みてtrueとする
		// https://gorm.io/docs/performance.html#Disable-Default-Transaction
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db.Logger = db.Logger.LogMode(logger.Silent)

	sqlDB, err := db.DB()
	// SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定します
	sqlDB.SetMaxIdleConns(DB_MAX_IDLE_CONN)
	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定します
	sqlDB.SetMaxOpenConns(DB_MAX_OPEN_CONN)
	// SetConnMaxLifetimeは再利用され得る最長時間を設定します
	sqlDB.SetConnMaxLifetime(DB_MAX_LIFETIME)
	if err != nil {
		panic(err)
	}
	return db
}

// 構造体に沿ったテーブルの作成を行う
func createTable(db *gorm.DB) {
	// テーブルの存在をチェックしない場合のみ作る
	if !db.Migrator().HasTable(&logMessage{}) {
		// テーブルの作成
		db.Migrator().CreateTable(&logMessage{})
	}
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
*/

type Account struct {
	AccountID int `gorm:"primaryKey"`
	Name      string
	Password  string
	CreatedAt time.Time
}

type Task struct {
	TaskID    int `gorm:"primaryKey"`
	Name      string
	Done      bool
	Message   string
	AccountID int `gorm:"foreignKey"`
	CreatedAt time.Time
}

func setTasktoList(task_id int, name string, msg string, done bool, account_id int) (r *Task) {
	r = new(Task)
	r.TaskID = task_id
	r.Name = name
	r.Message = msg
	r.Done = done
	r.AccountID = account_id
	return r
}

type TaskList = []*Task

func setupRouter() *gin.Engine {
	var tasks TaskList

	tasks = append(tasks, setTasktoList(1, "これはテストです1", "ああああああ", false, 1))
	tasks = append(tasks, setTasktoList(2, "これはテストです2", "ああああああ", true, 1))
	tasks = append(tasks, setTasktoList(3, "これはテストです3", "ああああああ", true, 1))

	r := gin.Default()
	r.Static("/static", "../static")
	r.LoadHTMLGlob("../templates/*.tmpl")

	fmt.Println(tasks)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.POST("/mypage", func(c *gin.Context) {
		var account = Account{
			Name:     c.PostForm("username"),
			Password: c.PostForm("password"),
		}

		c.HTML(http.StatusOK, "mypage.tmpl", gin.H{
			"username": account.Name,
			"tasks":    tasks,
		})
		//c.String(http.StatusOK, "home")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
