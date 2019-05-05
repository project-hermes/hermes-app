<template>
<main>
    <div class="inline-flex p-8 overflow-x-auto max-w-full">
        <h1 class="text-black">Dashboard</h1>
        <div aria-hidden="true" class="mx-12 border-l border-solid border-grey-lighter" />
        <DashboardAnalytics class="" />
    </div>
    <div class="h-82">
        <SimpleMap class="h-full" />
    </div>
    <div class="p-6 overflow-x-auto max-w-full">
        <DivesTable />
    </div>
</main>
</template>

<script>
import {SimpleMap} from '~/components';
import DashboardAnalytics from './DashboardAnalytics.vue';
import DivesTable from './DivesTable.vue';
export default {
    components: {
        SimpleMap,
        DashboardAnalytics,
        DivesTable
    }
};
</script>
<style>
    .border-left {

    }
</style>
