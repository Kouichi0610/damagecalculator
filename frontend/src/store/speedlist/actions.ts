import { ActionTree } from 'vuex';
import axios from 'axios';
import { SpeedListState, SpeedInfo } from './types';
import { RootState } from '../types';

class SpeedInfoImpl implements SpeedInfo {
  info: string;
  species: number;
  constructor(info: string, species: number) {
    this.info = info;
    this.species = species;
  }
}

export const actions: ActionTree<SpeedListState, RootState> = {
  getList: ({ commit }) => {
    axios.get('speed_list')
    .then((response) => {
      let json = JSON.stringify(response.data);
      let list = JSON.parse(json);

      let speeds: SpeedInfo[] = new Array();
      for (var i = 0; i < list.length; i++) {
        let s = new SpeedInfoImpl(list[i].info, list[i].speed);
        speeds.push(s)
      }
      commit('setInfos', speeds);
    })
    .catch((e) => {
      console.log('failed:' + e)
    });
    return true;
  }
}
