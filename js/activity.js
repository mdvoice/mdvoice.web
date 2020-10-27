$(function() {
  console.log(markdownit().render($("#md").html()));
  var md = markdownit({
    html: true,
    typographer: true
  }).use(markdownitEmoji)
  console.log(markdownit().render($("#md").html()));
  console.log($("#md").html());
  $("#md").html(md.render($("#md").html().replace(/&gt;/g,">").replace(/&lt;/g,"<")));
  // $("#md").html(markdownit().render($("#md pre").text()));
})

// function md() {
//   console.log(markdownit().render($("#md").html()));
//   $("#md").html();
//
// }
