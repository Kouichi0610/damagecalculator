// 名前入力 
// TODO:サジェスト機能入れたい
// $ npm install --save vue-cool-select
// TODO:picachuなど読み方指定したら行けるか

<template>
    <div class="names">
        <b-input-group size="sm">
            {{ tag }}
            <b-form-input placeholder="名前" list="tags-list" v-model="name"></b-form-input>
            <b-form-datalist id="tags-list" :options="candidates">
            </b-form-datalist>
        </b-input-group>

        <!-- 
        <input type="text" list="listId" v-model="name" placeholder="test">
        <datalist id="listId">
            <option
                v-for="(item, index) in candidates"
                :key="index"
                :value="item"
            ></option>
        </datalist>
        -->
        <!-- 
        <cool-select v-model="name" :items="candidates" placeholder="名前">
            <div slot="before-items-fixed" class="slot">Fixed header</div>
            <div slot="before-items" class="slot">Header</div>
            <div slot="after-items" class="slot slot--after">Footer</div>
            <div slot="after-items-fixed" class="slot slot--after">Fixed footer</div>
        </cool-select>
        -->
    </div>
</template>

<script>
//import {CoolSelect} from "vue-cool-select";

export default {
    name: 'Names',
    // 名前 設定済みの名前、候補一覧
    props: ['tag', 'name', 'candidates'],
    components: {
        //CoolSelect,
    },
    created: function() {
        console.log('Created:' + this.name);
        console.log('candidates' + this.candidates);

    },
    methods: {
        inputHandler: function($event) {
            console.log('Handler', $event.target.value);
        }
    }
}
</script>

<style scoped>
.slot {
    text-align: center;
    background: green;
}
.slot--after {
    background: blue;
    color: white;
}
</style>