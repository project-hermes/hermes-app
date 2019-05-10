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
                    type: 'spline',
                    animation: false
                },
                title: '',
                plotOptions: {
                    series: {
                        turboThreshold: 3000,
                        color: {
                            linearGradient: {x1: 0, x2: 0, y1: 0, y2: 1},
                            stops: [[0, '#ED5857'], [1, '#92B0FC']]
                        },
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
                    {
                        title: {
                            text: 'Temperature (°C)'
                        },
                        gridLineDashStyle: 'Dash',
                        gridLineColor: '#dddddf',
                        gridLineWidth: 1,
                        gridZIndex: 4
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
