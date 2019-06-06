<template>
    <main class="p-10">
        <NumberWidget class="m-4" label="Status" 
value="Active" color="green" />
        <NumberWidget
            class="m-4"
            label="Total Time Underwater"
            value="47h 22m"
        />
        <NumberWidget class="m-4" label="Firmware Version" 
value="1.06.09" />
    </main>
</template>
<script>
import {NumberWidget} from '~/components';

export default {
    components: {
        NumberWidget
    }
};
</script>
