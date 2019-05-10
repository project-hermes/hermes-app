<template>
    <LineChart :series="series" :chart="chart" 
class="" />
</template>

<script>
import {LineChart} from '~/components';
import times from 'lodash/times';
import get from 'lodash/get';

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
                    // type: 'spline',
                    plotBackgroundColor: {
                        linearGradient: {x1: 0, x2: 0, y1: 0, y2: 1},
                        stops: [
                            [0, 'rgba(77, 102, 202, 0.1)'], // start
                            [1, 'rgba(77, 102, 202, 0.8)'] // end
                        ]
                    },
                    animation: false
                },
                title: '',
                plotOptions: {
                    series: {
                        turboThreshold: 3000,
                        fillColor: '#fff',
                        color: '#4D66CA',
                        marker: {
                            enabled: false
                        },
                        animation: false,
                        tickmarkPlacement: 'on'
                    }
                },
                tooltip: {},
                xAxis: {
                    type: 'datetime',
                    dateTimeLabelFormats: {
                        minute: '%M:%S'
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
