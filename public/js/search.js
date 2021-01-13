$(() => {
  const url =
      "https://www.googleapis.com/customsearch/v1?key=AIzaSyCtW5TKhKc0mFufnI2qpvgq4xYkCsmOJ9k&cx=3c56dbb9081dd9b5c&q=",
    go = () => {
      if ($("#search-bar").val()) {
        $("#res").html(`
            <div class="text-center">
            <h5>搜尋中...</h5>
            </div>
            `);
        fetch(url + $("#search-bar").val())
          .then((res) => {
            return res.json();
          })
          .then(({ items }) => {
            if (!items) {
              $("#res").html(`
              <div class="text-center">
              <h5>無搜尋結果</h5>
              </div>
              `);
              return;
            }
            $("#res").html("");
            items.forEach(({ link, title }) => {
              let tmpl = `
                <li class="list-group-item">
                    <span class="badge bg text-light">${link
                      .replace(
                        "https://mdvoice.functionxyz.eu.org/html/activity/",
                        ""
                      )
                      .replace("/", "")
                      .replaceAll("%20", " ")}</span>
                    <a href="${link}" role="button" class=""> ${title} </a>
                </li>
              `;
              $("#res").append(tmpl);
            });
          });
      }
    };

  $("#search").on("click", go);
  $("#search-bar").on("keypress", (e) => {
    if (e.code == "Enter") {
      go();
    }
  });
});
