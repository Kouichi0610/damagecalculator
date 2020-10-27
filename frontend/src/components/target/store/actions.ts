import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types'
import { TargetState } from './types'

export const actions: ActionTree<TargetState, RootState> = {
  getSpecies: ({commit}, name: string) => {
    axios.get('get_species', {
      params: {
        Name: name,
      }
    })
    .then((response) => {
      let json = JSON.stringify(response.data);
      let sp = JSON.parse(json);
      console.log('success:' +  json);
      //commit('setCandidates', species);
      return true;
    })
    .catch((e) => {
        console.log('failed:' + e);
        return false;
    });
  }
}

