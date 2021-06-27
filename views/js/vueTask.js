var todoapp = new Vue({
  // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
  el: '#vue-app',

  // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
  data: {
    // 新規追加タスク
    newtask: {
      Title: ""
    },
    // タスク情報
    tasks: [],
    can_submit_search: false,
    options: {
      animation: 200
    },
  },

  // インスタンス作成時の処理
  created: function () {
    this.doFetchAlltasks()
  },

  // メソッド定義
  methods: {
    // Enterキーで送信可能にする
    enable_submit() {
      this.can_submit_search = true;
    },
    // 新規タスク入力時にEnterを押下した時の処理
    submit() {
      if (!this.can_submit_search) return;

      // タスクの追加
      this.doAddtask()

      this.can_submit_search = false;
    },

    // 全てのタスク情報を取得する
    doFetchAlltasks() {
      axios.get('/api/v1/task/list')
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            // サーバから取得したタスク情報をdataに設定する
            this.tasks = response.data.data
          }
        })
        .catch(error => {
          console.log(error);
          window.alert("ToDoリストの取得に失敗しました")
        })
      return
    },

    // タスク情報を登録する
    doAddtask() {

      // 更新値のチェック
      if (!this.ValidationCheck(this.newtask)) { return }

      // サーバへ送信するパラメータ
      const params = new URLSearchParams();
      params.append('Title', this.newtask.Title)
      // TODO ログイン実装時にやる
      params.append('UserID', 0)
      // TODO message実装時にやる
      params.append('Message', "")
      params.append('Done', false)

      axios.post('/api/v1/task/add', params)
        .then(response => {
          if (response.status != 201) {
            throw new Error('レスポンスエラー')
          } else {
            // タスク情報を取得する
            this.doFetchAlltasks()

            // 入力値を初期化する
            this.initInputValue()
          }
        })
        .catch(error => {
          console.log(error);
          window.alert("ToDoリストへの追加に失敗しました")
        })
    },

    // タスク情報を削除する
    doDeletetask(id) {
      axios.delete('/api/v1/task/delete/' + id)
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            // タスク情報を取得する
            this.doFetchAlltasks()
          }
        })
        .catch(error => {
          console.log(error);
          window.alert("ToDoリストからの削除に失敗しました")
        })
    },
    // タスク情報をアップデートする
    doUpdatetask(task) {

      // 更新値のチェック
      if (!this.ValidationCheck(task)) { return }

      // サーバへ送信するパラメータ
      const params = new URLSearchParams();
      params.append('Title', task.Title)
      // TODO ログイン実装時にやる
      params.append('UserID', 0)
      // TODO message実装時にやる
      params.append('Message', "")
      params.append('Done', task.Done)

      axios.put('/api/v1/task/update/' + task.TaskID, params)
        .then(response => {
          if (response.status != 201) {
            throw new Error('レスポンスエラー')
          } else {
            // タスク情報を取得する
            this.doFetchAlltasks()
          }
        })
        .catch(error => {
          console.log(error);
          window.alert("ToDoリストの更新に失敗しました")
        })
    },
    // タスクの順序をVMから取得する
    doFetchTaskVMOrder() {
      return this.tasks.map((e) => e.TaskID)
    },

    // タスクの順序をアップデートする
    doUpdateTaskOrder() {

      // 現在の配列の順序を取得
      taskOrderList = this.doFetchTaskVMOrder()

      // サーバへ送信するパラメータ
      const params = new URLSearchParams();
      params.append('OrderList', taskOrderList)
      // TODO ログイン実装時にやる
      params.append('UserID', 0)

      // APIを実行し順序を送信する
      axios.put('/api/v1/taskOrder/update', params)
        .then(response => {
          if (response.status != 201) {
            throw new Error('レスポンスエラー')
          } else {
            // タスク情報を取得する
            this.doFetchAlltasks()
          }
        })
        .catch(error => {
          console.log(error);
          window.alert("ToDoリストの更新に失敗しました")
        })
    },
    // 初期値を表示する
    initInputValue() {
      this.newtask = {
        Title: ""
      }
    },
    // 入力値のチェックを行う
    ValidationCheck(task) {
      if (task.Title.length < 2 || task.Title.length > 100) {
        window.alert("タスク名は2~100文字におさめてください")
        return false
      }
      return true
    },
    // タスクを入れ替えた時の発火処理
    draggableEnd(event) {

      // 順序変更のAPIを呼び出す
      this.doUpdateTaskOrder()
    }
  },
  mounted: async function () {
    // 5秒ごとにサーバのデータを取得する
    const intervalId = setInterval(await this.doFetchAlltasks, 5000)
  }
})