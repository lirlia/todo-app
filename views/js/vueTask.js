var todoapp = new Vue({
  // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
  el: '#vue-app',

  // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
  data: {
    // タスク情報
    tasks: []
  },

  // インスタンス作成時の処理
  created: function () {
    this.doFetchAlltasks()
  },

  // メソッド定義
  methods: {
    // 全てのタスク情報を取得する
    doFetchAlltasks() {
      axios.get('/api/v1/task/list')
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            var resulttasks = response.data

            // サーバから取得したタスク情報をdataに設定する
            this.tasks = resulttasks.data
            console.log(resulttasks)
          }
        })
    },

    // タスク情報を登録する
    /*
    doAddtask() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams();
      params.append('taskName', this.taskName)
      params.append('taskMemo', this.taskMemo)
  
      axios.post('/addtask', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            // タスク情報を取得する
            this.doFetchAlltasks()
  
            // 入力値を初期化する
            this.initInputValue()
          }
        })
    },
    */

    // タスク情報を削除する
    /*
    doDeletetask(task) {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams();
      params.append('taskID', task.id)
  
      axios.post('/deletetask', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            // タスク情報を取得する
            this.doFetchAlltasks()
          }
        })
    },
    */
  }
})