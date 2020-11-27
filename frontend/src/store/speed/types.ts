import axios from 'axios'
import { TargetCondition } from '../target/targetCondition'

export interface SpeedState {
  level: number;
  otherList: SpeedInfo[];
  targetSpeed: number; // とくせい、もちもの込みの素早さ
}

// すばやさ取得
export class GetSpeed {
  getSpeed(target: TargetCondition): Promise<number> {
    return new Promise((resolve, reject) => {
      axios.get('get_speed', {
        params: {
          Name: target.target,
          Level: target.level,
          Individuals: target.individuals,
          BasePoint: target.basePoints.speed(),
          Ability: target.ability,
          Nature: target.nature.name,
          Item: target.item.name,
          Condition: target.condition,
          Weather: target.weather.name,
          Field: target.field.name,
          }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let result = JSON.parse(json);
        resolve(result.speed)
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject(0);
      });
    });
  }
}

// 仮想敵　一覧 取得
export class SpeedList {
  speedList(level: number): Promise<SpeedInfo[]> {
    return new Promise((resolve, reject) => {
      axios.get('speed_list?Level=' + level)
      .then((response) => {
        let json = JSON.stringify(response.data);
        let list = JSON.parse(json);
        let res: SpeedInfo[] = [];
        for (var i = 0; i < list.length; i++) {
          res.push(new SpeedInfo(i, list[i].info, list[i].speed));
        }
        resolve(res);
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject([]);
      });
    });
  }

}

export class SpeedInfo {
  index: number;
  info: string;
  speed: number;

  constructor(index: number, info: string, speed: number) {
    this.index = index;
    this.info = info;
    this.speed = speed;
  }
}
