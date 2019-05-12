<template>
    <LineChart :series="series" :chart="chart" 
class="" />
</template>

<script>
import {LineChart} from '~/components';
import times from 'lodash/times';
import get from 'lodash/get';
import Highcharts from 'highcharts';

export default {
    components: {
        LineChart
    },
    data() {
        const d = Date.now();
        const data = times(10, i => {
            return {
                x: d + i * 60000,
                // y: i * 10
                y: (Math.random() * 100) | 0
            };
        });
        return {
            series: [{data}],
            chart: {
                chart: {
                    type: 'areaspline',
                    plotBackgroundColor: '#fff',
                    animation: false
                },
                title: '',
                plotOptions: {
                    series: {
                        turboThreshold: 3000,
                        fillColor: {
                            linearGradient: {x1: 0, x2: 0, y1: 0, y2: 1},
                            stops: [
                                [0, 'rgba(77, 102, 202, 0.1)'], // start
                                [1, 'rgba(77, 102, 202, 0.8)'] // end
                            ]
                        },
                        color: '#4D66CA',
                        marker: {
                            enabled: false
                        },
                        animation: false,
                        tickmarkPlacement: 'on'
                    }
                },
                tooltip: {
                    backgroundColor: '#333336',
                    padding: 10,
                    formatter() {
                        return `
                                <span style="padding: 10px;">
                                    <span style="color: #fff; font-weight: 600; font-size: 14px;">Depth:</span>
                                    <span style="color: #fff; font-size: 14px;">${
                                        this.y
                                    }m</span>
                                    <br />
                                    <span style="color: #fff; font-weight: 600; font-size: 14px;">Time:</span>
                                    <span style="color: #fff; font-size: 14px;">${Highcharts.dateFormat(
                                        '%H:%M:%S',
                                        new Date(+this.x)
                                    )}</span>
                                </span>
                        `;
                    }
                },
                xAxis: {
                    type: 'datetime',
                    dateTimeLabelFormats: {
                        minute: '%H:%M:%S'
                    },
                    title: {
                        text: 'Time'
                    },
                    min: get(data, '0.x'),
                    max: get(data, [get(data, ['length']) - 1, 'x'])
                },
                yAxis: [
                    // {
                    //     title: {
                    //         text: 'Temperature'
                    //     },
                    //     labels: {
                    //         format: '{value} °C'
                    //     }
                    // },
                    {
                        title: {
                            text: 'Depth (Meters)'
                        },
                        // labels: {
                        //     format: '{value}'
                        // },
                        reversed: true,
                        gridLineDashStyle: 'Dash',
                        gridLineColor: '#dddddf',
                        gridLineWidth: 1,
                        gridZIndex: 4
                        // opposite: true
                    }
                ],
                legend: {
                    enabled: false
                },
                credits: {
                    enabled: false
                }
            }
        };
    }
};
</script>
