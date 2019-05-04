<template>
    <router-link
        tag="div"
        :to="route"
        class="h-full px-1 uppercase text-xs flex items-center bg-blue hover:bg-blue-dark cursor-pointer sm:px-2 md:px-4 sm:text-sm"
        :class="{
            'text-white': isActive,
            'text-blue-light': !isActive
        }"
    >
        <slot />
    </router-link>
</template>
<script>
    export default {
        props: ['route']
    }
</script>
<style scoped>
.router-link-exact-active {
    color: #fff;
}
</style>
