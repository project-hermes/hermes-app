<template>
    <div :id="id" />
</template>

<script>
import Highcharts from 'highcharts';
import get from 'lodash/get';
export default {
    props: {
        series: {
            type: Array,
            default: () => []
        },
        chart: {
            type: Object,
            default: undefined
        }
    },
    data() {
        return {
            id: 'chart-' + Date.now() + Math.random()
        };
    },
    watch: {
        series() {
            this.renderChart();
        },
        chart() {
            this.renderChart();
        }
    },
    mounted() {
        this.id = this.renderChart();
    },
    methods: {
        renderChart() {
            if (!this.chart) return;

            Highcharts.chart(this.id, {
                ...this.chart,
                series: this.series
            });
        }
    }
};
</script>
