<template>
    <main>
        <div
            class="w-full flex items-center p-10 border-b border-grey-lighter overflow-x-auto"
        >
            <div class="pr-10  border-r border-grey-lighter">
                <h1 class="text-black font-normal text-4xl">Dashboard</h1>
            </div>
            <DashboardAnalytics class="pl-10" />
        </div>
        <div class="h-96">
            <SimpleMap class="h-full" />
        </div>
        <div class="p-8">
            <div class="">
                <div
                    class="border-t border-l border-r border-grey-lighter flex justify-between"
                >
                    <h3 class="font-normal p-4">Recent Dives</h3>
                    <router-link
                        to="/dives"
                        class="no-underline text-blue hover:underline p-4"
                    >
                        View all
                        <ChevronRightIcon class="h-3 w-3" />
                    </router-link>
                </div>
                <DivesTable :limit="5" />
            </div>
        </div>
    </main>
</template>

<script>
import {SimpleMap} from '~/components';
import DashboardAnalytics from './DashboardAnalytics.vue';
import DivesTable from '../DivesListView/DivesTable.vue';
import ChevronRightIcon from 'vue-feather-icons/icons/ChevronRightIcon';

export default {
    components: {
        SimpleMap,
        DashboardAnalytics,
        DivesTable,
        ChevronRightIcon
    }
};
</script>
<style>
.border-left {
}
</style>
