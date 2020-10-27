var offset;


$(function() {
  // var rellax = new Rellax('.rellax');
  offset = ($("#nav").css("width") == "0px") ? 0 : -60;
  $(window).resize(function(event) {
    offset = ($("#nav").css("width") == "0px") ? 0 : -60;
  });
  s();
  // paula=new paula();
  // $(window).scroll(function(event) {
  //   paula.scroll();
  // });
  // $(".navbtn").on('click',  function(event) {
  //   event.preventDefault();
  //   $("#nav").toggleClass(['animate__fadeIn','show']);
  // });
  $(".nav-link").on('click', function(event) {
    $("#open")[0].click();
  });
  // tks https://stackoverflow.com/a/37883208

  // $(window).resize(function(event) {
  // s();
  // });
});

function s() {
}

/*
var an=function () {
  var e=$(".an");
  e.find('.onclick').on('click', function(event) {
    event.preventDefault();
    $(this).attr('do');
    var r=new RegExp(/$t/);
    var i=()
    $($(this).attr('to')).
    ;
  });
}
*/
