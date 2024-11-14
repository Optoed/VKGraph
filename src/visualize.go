package src

import (
	"fmt"
	"net/http"
)

package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"net/http"
)

// Пример данных графа
var graph = map[int][]int{
	1: {2, 3},
	2: {1, 4, 5},
	3: {1, 6},
	4: {2},
	5: {2},
	6: {3},
}

// Функция для генерации графа
func generateGraphChart() *charts.Graph {
	// Инициализация графа с глобальными настройками
	graphChart := charts.NewGraph()
	graphChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "800px",
			Height: "600px",
		}),
		charts.WithTitleOpts(opts.Title{Title: "Визуализация Графа"}),
	)

	nodes := []opts.GraphNode{}
	links := []opts.GraphLink{}

	// Создаем узлы и связи
	for node, neighbors := range graph {
		nodes = append(nodes, opts.GraphNode{Name: fmt.Sprintf("%d", node)})
		for _, neighbor := range neighbors {
			links = append(links, opts.GraphLink{
				Source: fmt.Sprintf("%d", node),
				Target: fmt.Sprintf("%d", neighbor),
			})
		}
	}

	// Добавляем серию данных в граф
	graphChart.AddSeries("Graph", nodes, links).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Layout:    "force",
				Draggable: true,
				Roam:      true,
				Repulsion: 200, // Устанавливаем силу отталкивания между узлами
				Gravity:   0.1, // Настройка гравитации для удержания узлов ближе друг к другу
			}),
		)

	return graphChart
}

// HTTP-обработчик для рендера графа
func VisualizeGraphHandler(w http.ResponseWriter, _ *http.Request) {
	graphChart := generateGraphChart()
	w.Header().Set("Content-Type", "text/html")
	graphChart.Render(w)
}
