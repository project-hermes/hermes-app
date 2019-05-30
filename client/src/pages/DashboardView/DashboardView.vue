<template>
    <main class="h-full flex flex-col">
        <div
            class="h-32 z-10 w-full flex items-center p-10 border-b border-grey-lighter overflow-x-auto"
        >
            <div class="pr-10">
                <h1 class="text-black font-normal text-4xl">Dashboard</h1>
            </div>
            <DashboardAnalytics
                v-if="dives.length > 0"
                class="pl-10 border-l border-grey-lighter"
            />
        </div>
        <div v-if="dives.length > 0">
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
        </div>
        <div
            v-else
            class="bg-map flex-auto flex flex-col items-center justify-center h-full"
        >
            <h2 class="font-normal">
                You don't have any dives recorded in Project Hermes yet.
            </h2>
            <span class="text-grey-darker m-4">
                Sync data from your Hermes sensor to get started.
            </span>
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
    },
    data() {
        return {
            dives: [111]
        };
    }
};
</script>
<style>
.bg-map::after {
    content: '';
    position: absolute;
    top: 7.5rem;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: -1;
    background-image: url('../../img/empty-map-bg.png');
    background-repeat: no-repeat;
    background-size: cover;
    opacity: 0.3;
}
</style>
