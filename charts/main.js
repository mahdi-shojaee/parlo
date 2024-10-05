const { JSDOM } = require("jsdom");
const fs = require("node:fs");

const { window } = new JSDOM();
global.window = window;
global.document = window.document;
global.navigator = window.navigator;
global.navigator.userAgent = window.navigator.userAgent;
global.navigator.language = window.navigator.language;
const echarts = require("echarts");

echarts.setPlatformAPI({
  createCanvas() {},
});

function renderChart(data, outFilePath) {
  const fns = Object.keys(data.durations);

  const chart = echarts.init(null, null, {
    renderer: "svg",
    ssr: true,
    width: 500,
    height: 180,
  });

  const option = {
    animation: false,
    title: {
      text: data.title,
      top: "8%",
      left: "3%",
    },
    backgroundColor: "white",
    legend: {
      data: fns,
      top: "8%",
      right: "3%",
      align: "right",
    },
    grid: {
      left: "4%",
      right: "10%",
      bottom: "5%",
      containLabel: true,
    },
    tooltip: {},
    xAxis: {
      type: "category",
      boundaryGap: false,
      data: data.sizes,
    },
    yAxis: {
      type: "value",
    },
    series: Object.keys(data.durations).map((fnName) => ({
      name: fnName,
      type: "line",
      smooth: false,
      data: data.durations[fnName],
    })),
  };

  chart.setOption(option);
  const svgStr = chart.renderToSVGString();
  chart.dispose();

  return svgStr;
}

function sliceData(data, startIndex, endIndex = data.sizes.length) {
  return {
    ...data,
    sizes: data.sizes.slice(startIndex, endIndex),
    durations: Object.fromEntries(
      Object.entries(data.durations).map(([k, v]) => [
        k,
        v.slice(startIndex, endIndex),
      ])
    ),
  };
}

for (const funcName of process.argv.slice(2)) {
  const name = funcName.toLowerCase();
  const inFilePath = `./data/${name}.json`;
  const outStartFilePath = `./images/${name}-start.svg`;
  const outEndFilePath = `./images/${name}-end.svg`;

  const data = JSON.parse(fs.readFileSync(inFilePath, "utf8"));

  const startIndex =
    data.parallelThresholdSizesIndex <= 0
      ? data.sizes.length / 2
      : data.parallelThresholdSizesIndex + 1;
  const endIndex =
    data.parallelThresholdSizesIndex <= 0
      ? data.sizes.length / 2 + 1
      : data.parallelThresholdSizesIndex + 2;

  const startData = sliceData(data, 0, endIndex);
  const startSvgStr = renderChart(startData, outStartFilePath);
  fs.writeFileSync(outStartFilePath, startSvgStr);

  const endData = sliceData(data, startIndex);
  const endSvgStr = renderChart(endData, outEndFilePath);
  fs.writeFileSync(outEndFilePath, endSvgStr);
}
