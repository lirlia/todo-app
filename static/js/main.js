$(function() {

  $(".task-checkbox").change(function() {

    //console.log("hello")
    if ($(this).prop('checked')) {
      $(this).next().addClass("line-throungh")
      $(this).next().removeClass("text-decoration-none")
    } else {
      $(this).next().addClass("text-decoration-none")
      $(this).next().removeClass("line-throungh")
    }
  });
});