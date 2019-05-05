<template>
    <div>
        <label class="text-grey-darker text-sm leading-tight">{{label}}</label>
        <div v-if="value" class="text-blue text-2xl">{{value}}</div>
        <div v-else>---</div>
    </div>
</template>
<script>
    export default {
        props: {
            label: {
                type: String,
                required: true
            },
            value: {
                type: [String, Number]
            }
        }
    };
</script>
