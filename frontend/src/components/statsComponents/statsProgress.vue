// 能力値コンポーネント
<template>
    <div class="statsProgress">
        <b-progress class="mt-2" :max="max" show-value animated>
            <b-progress-bar :value="current" :label="label + current" variant="info"></b-progress-bar>
        </b-progress>
    </div>
</template>

<script>
export default {
    name: 'stats',
    props: ['label', 'current'],
    data: function() {
        return {
            max: 255,
            lv: 'abcd',
        }
    }
}
</script>
