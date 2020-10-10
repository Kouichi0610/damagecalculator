// 種族値コンポーネント
<template>
    <div class="species">
        <p>種族値</p>
        <div v-for="param in speciesValues" :key="param.key" class="row mb-1">
            <div class="col-sm-1">{{ param.label }}{{ param.value }}</div>
            <div class="col-sm-6 pt-1">
              <b-progress :max="255" show-value animated>
                <b-progress-bar :value="param.value" variant="warning"></b-progress-bar>
              </b-progress>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'species',
    props: ['params'],
    computed: {
        speciesValues: function() {
            let res = [];
            res.push({key: 0, label:'', value:this.params[0]});
            res.push({key: 1, label:'', value:this.params[1]});
            res.push({key: 2, label:'', value:this.params[2]});
            res.push({key: 3, label:'', value:this.params[3]});
            res.push({key: 4, label:'', value:this.params[4]});
            res.push({key: 5, label:'', value:this.params[5]});
            return res;
        }
    }
}
</script>

<style scoped>
</style>