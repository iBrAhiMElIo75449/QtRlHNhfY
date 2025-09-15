// 代码生成时间: 2025-09-15 18:28:29
package main

import (
    "buffalo"
    "github.com/markbates/buffalo/render"
    "log"
)

// ChartData 用于存储图表的数据
type ChartData struct {
    Labels []string `json:"labels"`
    Values []float64 `json:"values"`
}

// ChartRenderer 用于渲染图表的响应
type ChartRenderer struct {
    *render.GiraffeRenderer
}

// Render 渲染图表数据
func (cr ChartRenderer) Render(ctx buffalo.Context) error {
    c := ctx.Value("data").(*ChartData)
    return cr.Giraffe(ctx, c, "chart.html")
}

func main() {
    app := buffalo.Automatic(buffalo.Options{
        PreWares: []buffalo.PreWare{
            render.Provide(render.Options{
                Engine: &render.GiraffeEngine{
                    Extensions: []string{".html"},
                },
            })},
    })

    // 定义交互式图表生成器的路由
    app.POST("/generate-chart", func(c buffalo.Context) error {
        // 从请求中解析图表数据
        var chartData ChartData
        if err := c.Bind(&chartData); err != nil {
            return err
        }

        // 添加图表数据到上下文中，以便渲染
        c.Set("data", &chartData)

        // 设置渲染器
        return c.Render(200, ChartRenderer{})
    })

    // 启动BUFFALO应用
    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

// chart.html 用于生成交互式图表的HTML文件
// 请将此HTML文件放在 buffalo 项目的 templates 目录下，并根据需要进行编辑
// 例如:
// <html>
// <head>
//     <title>Interactive Chart</title>
//     <!-- 引入图表库，例如 Chart.js -->
// </head>
// <body>
//     <canvas id="myChart"></canvas>
//     <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
//     <script>
//         var ctx = document.getElementById("myChart").getContext("2d\);
//         var myChart = new Chart(ctx, {
//             type: 'bar',
//             data: {
//                 labels: {{ .Labels | js }},
//                 datasets: [{
//                     label: 'Demo Chart',
//                     data: {{ .Values | js }},
//                     backgroundColor:
//                     [
//                         'rgba(255, 99, 132, 0.2)',
//                         'rgba(54, 162, 235, 0.2)',
//                         // ...
//                     ],
//                     borderColor:
//                     [
//                         'rgba(255, 99, 132, 1)',
//                         'rgba(54, 162, 235, 1)',
//                         // ...
//                     ],
//                     borderWidth: 1
//                 }]
//             },
//             options: {
//                 scales: {
//                     y: {
//                         beginAtZero: true
//                     }
//                 }
//             }
//         });
//     </script>
// </body>
// </html>