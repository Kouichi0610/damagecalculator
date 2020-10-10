// 性格変更コンポーネント
<template>
    <div class="nature">
        <b-form-select v-model="selected" :options="natures" size="sm" class="mt-3"></b-form-select>
    </div>
</template>

<script>
export default {
    name: 'nature',
    props: ['current', 'natures'],
    data: function() {
        return {
            selected: 'てれや'
        }
    },
    created: function() {
        this.selected = this.current;
    },
    watch: {
        selected: function() {
            this.$emit('update', this.selected);
        }
    }

}
</script>

<style scoped>
</style>