<template>
    <main class="h-full w-full">
        <div class="w-full flex items-center p-10 border-b border-grey-lighter">
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
        <div class="p-8" />
    </main>
</template>
<script>
import ChevronLeftIcon from 'vue-feather-icons/icons/ChevronLeftIcon';
import DiveNumbers from './DiveNumbers.vue';
export default {
    components: {
        ChevronLeftIcon,
        DiveNumbers
    }
};
</script>
<style>
.transform {
    transform: translate(0, 1px);
}
</style>
