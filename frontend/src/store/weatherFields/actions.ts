import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { WeatherFieldsState } from './types'

export const actions: ActionTree<WeatherFieldsState, RootState> = {
  getWeatherFields: ({ commit } ) => {
    axios.get('weather_fields')
    .then((response) => {
      let json = JSON.stringify(response.data);
      let res = JSON.parse(json);

      commit('setWeathers', res.weathers);
      commit('setFields', res.fields);
    })
    .catch((e) => {
      console.log('failed:' + e);
    });
  }
}
