// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

export async function search(params: API.SearchParam) {
    return request<API.ServiceResponse>('/service/search/stocks', {
      method: 'POST',
      data: params,
    });
  }

  export async function saveCustomCondition(params: any) {
    return request<API.ServiceResponse>('/service/common/saveCustomCondition', {
      method: 'POST',
      data: params,
    });
  }
  
  export async function deleteCondition(params: any) {
    return request<API.ServiceResponse>('/service/common/deleteCondition', {
      method: 'POST',
      data: params,
    });
  }
  
  export async function updateStockAttr(params: any) {
    return request<API.ServiceResponse>('/service/update/updateStockAttr', {
      method: 'POST',
      data: params,
    });
  }

  export async function blackList(params: any) {
    return request<API.ServiceResponse>('/service/update/blackList', {
      method: 'POST',
      data: params,
    });
  }
  
  export async function fetchCustomConditionName() {
    return request<API.ServiceResponse>(`/service/common/fetchCustomConditionName`);
  }
  
  export async function fetchCustomConditionByName(params: any) {
    return request<API.ServiceResponse>('/service/common/fetchCustomConditionByName', {
      method: 'POST',
      data: params,
    });
  }
  
  export async function getHistoryBySymbol(params: any) {
    return request<API.ServiceResponse>('/service/kline/getHistoryBySymbol', {
      method: 'POST',
      data: params,
    });
  }