// TODO:
import {mount} from '@vue/test-utils'
import { shallowMount } from '@vue/test-utils'
import {sum} from './sum'

describe('sum sample', (): void => {
  test('1+1 = 3', (): void => {
    let response: number = sum(1, 5);
    expect(response).toBe(6);
  });
})

