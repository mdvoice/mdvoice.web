var paula;
$(function() {
  AOS.init({
    offset: 300
  });
  $('a[href*="#"]:not([href="#"]).scroll').click(function() {
    if (location.pathname.replace(/^\//, '') == this.pathname.replace(/^\//, '') && location.hostname == this.hostname) {
      var target = $(this.hash);
      target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
      if (target.length) {
        $('html, body').animate({
          scrollTop: target.offset().top + offset
        }, 500);
        return false;
      }
    }
  });

})
