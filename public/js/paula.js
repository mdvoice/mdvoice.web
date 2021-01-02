/**
 * @copyright made my kaiyo hugo
 */






/**
 * var paula - init paula
 *
 * @param  {type} a array for config
 * @param  {type} a.class set class
 * @param  {type} a.offset set offset
 *
 * @return {type} true
 */
var paula = function(a) {
  this.select = (a && a.select) ? a.select : ".paula",
    this.offset = (a && a.offset) ? a.offset : 0,
    this.range = (a && a.range) ? a.range : 0,
    this.mode = (a && a.mode) ? a.mode : 0,
    this.element = $(this.select),
    this.reflesh = function() {
      this.element = $(this.select);
    },
    this.scroll = function() {
      var scroll = $(window).scrollTop(),
        viewport = $(window).height(),
        offset = this.offset,
        range = this.range,
        mode = this.mode,
        element = this.element;
      element.each(function() {
        var show = scroll + viewport - $(this).offset().top - offset,
          canwatch = (show >= 0 && show <= viewport + range + $(this).offset().top || mode) ? true : false;
        console.log(show);
        console.log(canwatch);
        // console.log($(this).css('animate-delay'));
        // for DEBUG
        if (canwatch) {
          $(this).css('animation-delay', "-" + ($(this).css('animation-duration').replace(/s$/, "") * (show / (viewport + range)) + "s"));
        }
      });
    };
  this.element.css('animation-play-state', 'paused');
  return true;
}
