// jquery & js for bootstrap
require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");


$(() => {
  console.log('we here');


  let heatmap = new Chart({
    parent: "#chart",
    type: "heatmap",
    height: 115,
    data: heatmap_data, // object with date/timestamp-value pairs

    discrete_domains: 1, // default: 0

    start: start_date,
    // A Date object;
    // default: today's date in past year
    //          for an annual heatmap

    legend_colors: ["#ebedf0", "#fdf436", "#ffc700", "#ff9100", "#06001c"]
    // Set of five incremental colors,
    // beginning with a low-saturation color for zero data;
    // default: ['#ebedf0', '#c6e48b', '#7bc96f', '#239a3b', '#196127']
  });
}); // end of js