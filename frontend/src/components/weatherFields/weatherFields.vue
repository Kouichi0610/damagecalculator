<template>
  <div class="WeatherFields">
    <div class="row mb-1">
      <div class="col-1"></div>
      <b-dropdown class="col-1" id="weather-dropdown" text="天候">
        <b-dropdown-item v-for="item in weathers" :key="item" @click="changeWeather(item)">{{ item }}</b-dropdown-item>
      </b-dropdown>
      <b-dropdown class="col-1" id="field-dropdown" text="フィールド">
        <b-dropdown-item v-for="item in fields" :key="item" @click="changeField(item)">{{ item }}</b-dropdown-item>
      </b-dropdown>
      <div class="col-1">{{current}}</div>
    </div>
  </div>
</template>

<script lang="ts">
/*
  天候、フィールド
*/
import { Vue, Prop, Component, Watch } from "vue-property-decorator";
import { Getter, Action } from 'vuex-class';
import { WeatherFieldsState } from '../../store/weatherFields/types';

const namespace: string = 'weatherFields';

@Component
export default class WeatherFields extends Vue {
  @Action('getWeatherFields', { namespace })
  private getWeatherFields!: () => void;
  @Getter('isInitialized', { namespace })
  private isInitialized!: boolean;
  @Getter('weathers', { namespace })
  private weathers!: string[];
  @Getter('fields', { namespace })
  private fields!: string[];


  created() {
    this.getWeatherFields();
  }

  changeWeather(weather: string) {
    console.log('Weather:' + weather);
  }
  changeField(field: string) {
    console.log('Field:' + field);
  }
  
}
</script>

<style scoped>
</style>
