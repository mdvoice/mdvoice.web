$(function() {
  AOS.init();
  $(".navbtn").on('click',  function(event) {
    event.preventDefault();
    $("#nav").toggleClass(['animate__fadeIn','show']);
  });
  $(".nav-link").on('click', function(event) {
    $("#nav").removeClass(['animate__fadeIn','show']);
  });
});


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
