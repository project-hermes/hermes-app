<template>
    <div class="inline-flex">
        <NumberWidget
            v-for="widget in widgets"
            class="mr-10"
            :key="widget.label"
            :label="widget.label"
            :value="widget.value"
        />
    </div>
</template>
<script>
    import NumberWidget from './NumberWidget.vue';
    export default {
        components: {
            NumberWidget
        },
        data () {
            return {
                widgets: [
                    {
                        label: 'Dives',
                        value: 23
                    },
                    {
                        label: 'Total Time',
                        value: '47h 22m'
                    },
                    {
                        label: 'Max Depth',
                        value: '31.24m'
                    },
                    {
                        label: 'Temp High',
                        value: '88˚F'
                    },
                    {
                        label: 'Temp Low',
                        value: '57˚F'
                    }
                ]
            }
        }
    };
</script>
