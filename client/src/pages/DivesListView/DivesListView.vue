<template>
    <main class="h-full w-full">
        <div class="w-full p-10 inline-flex items-center items-stretch border-b border-grey-lighter">
            <h1 class="font-normal text-4xl">
                Dives
                <span class="text-grey-darker text-2xl mx-2">
                    (82)
                </span>
            </h1>
        </div>
        <div class="p-8">
            <DivesTable />
        </div>
    </main>
</template>
<script>
    import DivesTable from './DivesTable.vue';
    export default {
        components: {
            DivesTable
        }
    };
</script>
