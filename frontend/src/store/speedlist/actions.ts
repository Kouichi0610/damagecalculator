import { ActionTree } from 'vuex';
import axios from 'axios';
import { SpeedListState, SpeedInfo } from './types';
import { RootState } from '../types';

// TODO:何処かstoreに
const Level = 50;

class SpeedInfoImpl implements SpeedInfo {
  info: string;
  speed: number;
  constructor(info: string, speed: number) {
    this.info = info;
    this.speed = speed;
  }
}

export const actions: ActionTree<SpeedListState, RootState> = {
  getList: ({ commit }) => {
    axios.get('speed_list?Level=' + Level)
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
