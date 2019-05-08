<template>
    <router-link
        :to="route"
        class="h-full px-1 uppercase no-underline text-xs text-blue-light flex items-center bg-blue hover:bg-blue-dark cursor-pointer sm:px-2 md:px-4 sm:text-sm md:mx-1"
    >
        <slot />
    </router-link>
</template>
<script>
export default {
    props: ['route']
};
</script>
<style scoped>
.router-link-active {
    color: #fff;
}
</style>
