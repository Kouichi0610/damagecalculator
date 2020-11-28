import axios from 'axios'
import { TargetCondition } from '../target/targetCondition'

export class SpeedOrderOld {
  speedList(targetCondition: TargetCondition) {
    axios.get('speed_list?Level=' + targetCondition.level)
    .then((response) => {
      let json = JSON.stringify(response.data);
      let list = JSON.parse(json);

      console.log('SpeedList' + list.length);
      for (var i = 0; i < list.length; i++) {
        console.log('' + list[i].info + ' Speed:' + list[i].speed);
      }
    })
    .catch((e) => {
      console.log('failed:' + e)
    });
  }
  // 特性が自身のすばやさに与える影響
  abilityOwner(targetCondition: TargetCondition) {
    axios.get('ability_owner_speed?Ability=' + targetCondition.ability)
    .then((response) => {
      let json = JSON.stringify(response.data);
      let a = JSON.parse(json);
      console.log('自身に:' + a.comment + ' ランク補正:' + a.rank + ' 倍率:' + a.magnification);
    })
    .catch((e) => {
      console.log('failed:' + e)
    });
  }
  // 特性が相手のすばやさに与える影響
  abilityOther(targetCondition: TargetCondition) {
    axios.get('ability_other_speed?Ability=' + targetCondition.ability)
    .then((response) => {
      let json = JSON.stringify(response.data);
      let a = JSON.parse(json);
      console.log('相手に:' + a.comment + ' ランク補正:' + a.rank + ' 倍率:' + a.magnification);
    })
    .catch((e) => {
      console.log('failed:' + e)
    });
  }

}