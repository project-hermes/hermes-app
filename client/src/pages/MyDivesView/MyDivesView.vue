<template>
<main>
    <div class="inline-flex m-8">
        <h1 class="text-black">My Dives</h1>
        <div aria-hidden="true" class="mx-12 border-l border-solid border-grey-lighter" />
        <MyDiveAnalytics class="" />
    </div>
    <div class="h-82">
        <SimpleMap class="h-full" />
    </div>
    <div class="p-6">
        <DivesTable />
    </div>
</main>
<!--  header -->
<!--  nums -->
<!--  map -->
<!-- table -->

</template>

<script>
import {SimpleMap} from '~/components';
import MyDiveAnalytics from './MyDiveAnalytics.vue';
import DivesTable from './DivesTable.vue';
export default {
    components: {
        SimpleMap,
        MyDiveAnalytics,
        DivesTable
    }
};
</script>
<style>
    .border-left {

    }
</style>
