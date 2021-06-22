$(function () {

  // タスク追加ボタン押下時の処理
  $("#add-task").on('click', function () {

    // validation check
    // 文字数が1文字未満の場合
    if ($("#add-task-value").val().length < 1) {
      window.alert("タスクは1文字以上で入力してください");
      return
    }

    // TODOリストに追加する
    $("#task-list").prepend(`
      <div class="panel-block columns py-1">
      <div class="checkbox column is-1 py-0 has-text-centered">
        <input class="task-checkbox" type="checkbox">
      </div>
      <div class="column py-0">
        ${$("#add-task-value").val()}
      </div>
      <div class="column is-1 py-0">
        <button class="delete-task button is-danger is-outlined is-small">
          <span>削除</span>
          <span class="icon is-small">
            <i class="fas fa-times"></i>
          </span>
        </button>
      </div>
    `);
    // 入力したタスク名をクリアする
    $("#add-task-value").val("")

    // TODO: DBの更新を行う
  });

  // タスク追加ボタン押下時の処理
  $(document).on('click', ".delete-task", function () {
    if (!confirm('本当に削除しますか？')) {
      // キャンセルの時の処理
      return false;
    } else {
      // OKの時の処理
      $(this).parent().parent().remove();

      // TODO: DB更新
    }
  });

  // チェックボックス押下時に取り消し線を入れる
  $(document).on('change', ".task-checkbox", function () {
    if ($(this).prop('checked')) {
      $(this).parent().next().addClass("line-throungh")
      $(this).parent().next().removeClass("text-decoration-none")
    } else {
      $(this).parent().next().addClass("text-decoration-none")
      $(this).parent().next().removeClass("line-throungh")
    }
  });
});