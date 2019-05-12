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
import {NumberWidget} from '~/components';
export default {
    components: {
        NumberWidget
    },
    data() {
        return {
            widgets: [
                {
                    label: 'Dives',
                    value: 23
                },
                {
                    label: 'Total Time',
                    value: '47h 22m'
                }
            ]
        };
    }
};
</script>
