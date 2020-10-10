// 能力値コンポーネント
<template>
    <div class="stats">
        <p>能力値</p>
        <div v-for="param in statsValues" :key="param.key" class="row mb-1">
            <div class="col-sm-3">{{ param.label }}{{ param.value }}</div>
            <div class="col-sm-8 pt-1">
              <b-progress :max="255" show-value animated>
                <b-progress-bar :value="param.value" variant="info"></b-progress-bar>
              </b-progress>
            </div>
        </div>
    </div>
</template>

<script>
//import StatsProgress from './statsProgress'

export default {
    name: 'stats',
    components: {
        //'stats-progress': StatsProgress,
    },
    props: ['params'],
    computed: {
        statsValues: function() {
            let res = [];
            res.push({key: 0, label:'HP:', value:this.params[0]});
            res.push({key: 1, label:'攻撃:', value:this.params[1]});
            res.push({key: 2, label:'防御:', value:this.params[2]});
            res.push({key: 3, label:'特攻:', value:this.params[3]});
            res.push({key: 4, label:'特防:', value:this.params[4]});
            res.push({key: 5, label:'素早さ:', value:this.params[5]});
            return res;
        }

    },
}
</script>

<style scoped>
.stats {
    width: 500px;
}
</style>