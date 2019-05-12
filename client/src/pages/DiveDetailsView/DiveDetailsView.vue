<template>
    <main class="h-full w-full">
        <div
            class="w-full flex items-center p-10 border-b border-grey-lighter overflow-x-auto"
        >
            <div class="pr-10 border-r border-grey-lighter">
                <router-link
                    :to="{name: 'dives'}"
                    class="no-underline text-blue hover:underline"
                >
                    <ChevronLeftIcon class="h-4 w-4 transform" />
                    Back to Dives
                </router-link>
                <h1 class="font-normal text-4xl py-3">
                    09/24/2018 09:43
                </h1>
                <span class="text-grey-darker">
                    37.780009, 23.336231 (Greece)
                </span>
            </div>
            <DiveNumbers class="pl-10" />
        </div>
        <div class="p-8">
            <div class="border border-grey-lighter mb-8">
                <div class="border-b border-grey-lighter">
                    <h3 class="font-normal p-4">Depth</h3>
                </div>
                <div class="flex">
                    <div
                        class="border-r border-grey-lighter flex-shrink flex flex-col"
                    >
                        <NumberWidget
                            class="p-6"
                            label="Max Depth"
                            value="31.24m"
                        />
                        <NumberWidget
                            class="p-6 pt-0"
                            label="Avg Depth"
                            value="24.51m"
                        />
                    </div>
                    <div class="w-full p-6 overflow-x-auto">
                        <DepthChart />
                    </div>
                </div>
            </div>
            <div class="border border-grey-lighter">
                <div class="border-b border-grey-lighter">
                    <h3 class="font-normal p-4">Temperature</h3>
                </div>
                <div class="flex">
                    <div
                        class="border-r border-grey-lighter flex-shrink flex flex-col"
                    >
                        <NumberWidget
                            class="p-6"
                            label="Max Temp"
                            value="27ºC"
                            color="red"
                        />
                        <NumberWidget
                            class="p-6 pt-0"
                            label="Avg Temp"
                            value="25ºC"
                            color="black"
                        />
                        <NumberWidget
                            class="p-6 pt-0"
                            label="Min Temp"
                            value="23ºC"
                            color="blue-light"
                        />
                    </div>
                    <div class="w-full p-6 overflow-x-auto">
                        <TemperatureChart />
                    </div>
                </div>
            </div>
        </div>
    </main>
</template>
<script>
import ChevronLeftIcon from 'vue-feather-icons/icons/ChevronLeftIcon';
import DiveNumbers from './DiveNumbers.vue';
import DepthChart from './DepthChart.vue';
import TemperatureChart from './TemperatureChart.vue';
import {NumberWidget} from '~/components';

export default {
    components: {
        ChevronLeftIcon,
        DiveNumbers,
        DepthChart,
        TemperatureChart,
        NumberWidget
    }
};
</script>
<style>
.transform {
    transform: translate(0, 1px);
}
</style>
