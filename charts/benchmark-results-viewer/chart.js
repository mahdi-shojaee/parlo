function renderChart(data, chartDom, numCore) {
  const fns = getFunctionNames(data);

  const chart = echarts.init(chartDom, null, {
    renderer: "svg",
  });

  const option = {
    backgroundColor: "#fff",
    title: {
      text: `GOMAXPROCS=${numCore}`,
      top: "5%",
      left: "2%",
      textStyle: {
        fontSize: 14,
      },
      subtextStyle: {
        fontSize: 14,
      },
    },
    legend: {
      data: fns,
      top: "5%",
      right: "2%",
      textStyle: {
        fontSize: 14,
      },
    },
    toolbox: {
      show: true,
      orient: "vertical",
      top: "center",
      right: "2%",
      feature: {
        dataZoom: {
          yAxisIndex: "none",
        },
        magicType: { type: ["line", "bar"] },
        restore: {},
      },
    },
    grid: {
      top: "25%",
      left: "17%",
      right: "15%",
      bottom: "25%",
      containLabel: false,
    },
    tooltip: {
      trigger: "axis",
      textStyle: {
        fontSize: 14,
      },
    },
    xAxis: {
      type: "category",
      boundaryGap: false,
      name: "Input Size",
      nameLocation: "end",
      nameGap: 10,
      nameTextStyle: {
        fontSize: 14,
      },
      axisLabel: {
        interval: 0,
        align: "left",
        padding: [5, 0, 0, 0],
        fontSize: 14,
      },
      data: data.sizes.map((size, index) => ({
        value: size.toLocaleString(),
        textStyle: {
          color:
            data.durations[fns[0]][numCore][index] >
            data.durations[fns[1]][numCore][index]
              ? "#008000"
              : "#FF0000",
          fontSize: 14,
        },
      })),
    },
    yAxis: {
      type: "value",
      name: "Duration (µs/op)",
      nameLocation: "end",
      nameGap: 20,
      scale: true,
      nameTextStyle: {
        fontSize: 14,
      },
      axisLabel: {
        fontSize: 14,
      },
    },
    series: fns.map((fnName, index) => {
      if (index === 0) {
        return {
          name: fnName,
          type: "line",
          data: data.durations[fnName][numCore].map((value) =>
            Math.floor(value / 1000)
          ),
        };
      }
      return {
        name: fnName,
        type: "line",
        data: data.durations[fnName][numCore].map((value, i) => {
          const oldValue = data.durations[fns[0]][numCore][i];
          if (oldValue === 0) {
            return {
              value: Math.floor(value / 1000),
              label: {
                show: false,
              },
            };
          }
          const improvement = ((oldValue - value) / oldValue) * 100;
          const speedup = oldValue / value;
          return {
            value: Math.floor(value / 1000),
            label: {
              show: true,
              position: "top",
              formatter: `${improvement.toFixed(1)}%\n${speedup.toFixed(2)}x`,
              fontSize: 14,
              fontWeight: "bold",
              padding: [4, 4],
              borderRadius: 4,
              backgroundColor:
                improvement > 0
                  ? "rgba(0, 128, 0, 0.5)"
                  : "rgba(255, 0, 0, 0.5)",
              color: improvement > 0 ? "#E0FFE0" : "#FFEEEE",
              lineHeight: 14,
            },
          };
        }),
      };
    }),
    dataZoom: [
      {
        type: "slider",
        xAxisIndex: 0,
        filterMode: "filter",
        height: 16,
        bottom: 36,
        borderColor: "transparent",
        backgroundColor: "#f0f0f0",
        fillerColor: "rgba(200, 200, 200, 0.05)",
        handleStyle: {
          color: "#a0c8ff",
        },
        throttle: 0,
        textStyle: {
          fontSize: 14,
        },
      },
    ],
  };

  function updateChartForScreenSize() {
    const isSmallScreen = window.innerWidth < 1024;
    chart.setOption({
      xAxis: {
        axisLabel: {
          rotate: isSmallScreen ? 20 : 0,
        },
      },
    });
  }

  // Set initial options and update for screen size
  chart.setOption(option);
  updateChartForScreenSize();

  // Add resize event listeners
  window.addEventListener("resize", () => {
    chart.resize();
    updateChartForScreenSize();
  });

  // Initial resize to fit the container
  chart.resize();
}

function isParFunction(functionName) {
  // Regular expression to match the pattern
  const regex = /^(?:parlo\.)?Par[A-Z]/;
  return regex.test(functionName);
}

function getFunctionNames(data) {
  return Object.keys(data.durations).sort((a, _) => {
    return isParFunction(a) ? 1 : -1;
  });
}
