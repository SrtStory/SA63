/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Statistic
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    EntArea,
    EntAreaFromJSON,
    EntAreaToJSON,
    EntCarrier,
    EntCarrierFromJSON,
    EntCarrierToJSON,
    EntContagious,
    EntContagiousFromJSON,
    EntContagiousToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientToJSON,
    EntStatistic,
    EntStatisticFromJSON,
    EntStatisticToJSON,
} from '../models';

export interface CreateAreaRequest {
    area: EntArea;
}

export interface CreateCarrierRequest {
    carrier: EntCarrier;
}

export interface CreateContagiousRequest {
    contagious: EntContagious;
}

export interface CreateEmployeeRequest {
    employee: EntEmployee;
}

export interface CreatePatientRequest {
    patient: EntPatient;
}

export interface CreateStatisticRequest {
    statistic: EntStatistic;
}

export interface GetAreaRequest {
    id: number;
}

export interface GetCarrierRequest {
    id: number;
}

export interface GetContagiousRequest {
    id: number;
}

export interface GetEmployeeRequest {
    id: number;
}

export interface GetPatientRequest {
    id: number;
}

export interface ListAreaRequest {
    limit?: number;
    offset?: number;
}

export interface ListCarrierRequest {
    limit?: number;
    offset?: number;
}

export interface ListContagiousRequest {
    limit?: number;
    offset?: number;
}

export interface ListEmployeeRequest {
    limit?: number;
    offset?: number;
}

export interface ListPatientRequest {
    limit?: number;
    offset?: number;
}

export interface ListStatisticRequest {
    limit?: number;
    offset?: number;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create area
     * Create area
     */
    async createAreaRaw(requestParameters: CreateAreaRequest): Promise<runtime.ApiResponse<EntArea>> {
        if (requestParameters.area === null || requestParameters.area === undefined) {
            throw new runtime.RequiredError('area','Required parameter requestParameters.area was null or undefined when calling createArea.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/areas`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntAreaToJSON(requestParameters.area),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntAreaFromJSON(jsonValue));
    }

    /**
     * Create area
     * Create area
     */
    async createArea(requestParameters: CreateAreaRequest): Promise<EntArea> {
        const response = await this.createAreaRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create carrier
     * Create carrier
     */
    async createCarrierRaw(requestParameters: CreateCarrierRequest): Promise<runtime.ApiResponse<EntCarrier>> {
        if (requestParameters.carrier === null || requestParameters.carrier === undefined) {
            throw new runtime.RequiredError('carrier','Required parameter requestParameters.carrier was null or undefined when calling createCarrier.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/carriers`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntCarrierToJSON(requestParameters.carrier),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCarrierFromJSON(jsonValue));
    }

    /**
     * Create carrier
     * Create carrier
     */
    async createCarrier(requestParameters: CreateCarrierRequest): Promise<EntCarrier> {
        const response = await this.createCarrierRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create contagious
     * Create contagious
     */
    async createContagiousRaw(requestParameters: CreateContagiousRequest): Promise<runtime.ApiResponse<EntContagious>> {
        if (requestParameters.contagious === null || requestParameters.contagious === undefined) {
            throw new runtime.RequiredError('contagious','Required parameter requestParameters.contagious was null or undefined when calling createContagious.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/contagiouss`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntContagiousToJSON(requestParameters.contagious),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntContagiousFromJSON(jsonValue));
    }

    /**
     * Create contagious
     * Create contagious
     */
    async createContagious(requestParameters: CreateContagiousRequest): Promise<EntContagious> {
        const response = await this.createContagiousRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create employee
     * Create employee
     */
    async createEmployeeRaw(requestParameters: CreateEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.employee === null || requestParameters.employee === undefined) {
            throw new runtime.RequiredError('employee','Required parameter requestParameters.employee was null or undefined when calling createEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/Employees`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntEmployeeToJSON(requestParameters.employee),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * Create employee
     * Create employee
     */
    async createEmployee(requestParameters: CreateEmployeeRequest): Promise<EntEmployee> {
        const response = await this.createEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create patient
     * Create patient
     */
    async createPatientRaw(requestParameters: CreatePatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.patient === null || requestParameters.patient === undefined) {
            throw new runtime.RequiredError('patient','Required parameter requestParameters.patient was null or undefined when calling createPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/patients`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntPatientToJSON(requestParameters.patient),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * Create patient
     * Create patient
     */
    async createPatient(requestParameters: CreatePatientRequest): Promise<EntPatient> {
        const response = await this.createPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create statistic
     * Create statistic
     */
    async createStatisticRaw(requestParameters: CreateStatisticRequest): Promise<runtime.ApiResponse<EntStatistic>> {
        if (requestParameters.statistic === null || requestParameters.statistic === undefined) {
            throw new runtime.RequiredError('statistic','Required parameter requestParameters.statistic was null or undefined when calling createStatistic.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/statistic`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntStatisticToJSON(requestParameters.statistic),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntStatisticFromJSON(jsonValue));
    }

    /**
     * Create statistic
     * Create statistic
     */
    async createStatistic(requestParameters: CreateStatisticRequest): Promise<EntStatistic> {
        const response = await this.createStatisticRaw(requestParameters);
        return await response.value();
    }

    /**
     * get area by ID
     * Get a area entity by ID
     */
    async getAreaRaw(requestParameters: GetAreaRequest): Promise<runtime.ApiResponse<EntArea>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getArea.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/areas/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntAreaFromJSON(jsonValue));
    }

    /**
     * get area by ID
     * Get a area entity by ID
     */
    async getArea(requestParameters: GetAreaRequest): Promise<EntArea> {
        const response = await this.getAreaRaw(requestParameters);
        return await response.value();
    }

    /**
     * get carrier by ID
     * Get a carrier entity by ID
     */
    async getCarrierRaw(requestParameters: GetCarrierRequest): Promise<runtime.ApiResponse<EntCarrier>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getCarrier.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/carriers/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCarrierFromJSON(jsonValue));
    }

    /**
     * get carrier by ID
     * Get a carrier entity by ID
     */
    async getCarrier(requestParameters: GetCarrierRequest): Promise<EntCarrier> {
        const response = await this.getCarrierRaw(requestParameters);
        return await response.value();
    }

    /**
     * get contagious by ID
     * Get a contagious entity by ID
     */
    async getContagiousRaw(requestParameters: GetContagiousRequest): Promise<runtime.ApiResponse<EntContagious>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getContagious.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/contagiouss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntContagiousFromJSON(jsonValue));
    }

    /**
     * get contagious by ID
     * Get a contagious entity by ID
     */
    async getContagious(requestParameters: GetContagiousRequest): Promise<EntContagious> {
        const response = await this.getContagiousRaw(requestParameters);
        return await response.value();
    }

    /**
     * get employee by ID
     * Get a employee entity by ID
     */
    async getEmployeeRaw(requestParameters: GetEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * get employee by ID
     * Get a employee entity by ID
     */
    async getEmployee(requestParameters: GetEmployeeRequest): Promise<EntEmployee> {
        const response = await this.getEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get patient by ID
     * Get a patient entity by ID
     */
    async getPatientRaw(requestParameters: GetPatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * get patient by ID
     * Get a patient entity by ID
     */
    async getPatient(requestParameters: GetPatientRequest): Promise<EntPatient> {
        const response = await this.getPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * list area entities
     * List area entities
     */
    async listAreaRaw(requestParameters: ListAreaRequest): Promise<runtime.ApiResponse<Array<EntArea>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/areas`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntAreaFromJSON));
    }

    /**
     * list area entities
     * List area entities
     */
    async listArea(requestParameters: ListAreaRequest): Promise<Array<EntArea>> {
        const response = await this.listAreaRaw(requestParameters);
        return await response.value();
    }

    /**
     * list carrier entities
     * List carrier entities
     */
    async listCarrierRaw(requestParameters: ListCarrierRequest): Promise<runtime.ApiResponse<Array<EntCarrier>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/carriers`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntCarrierFromJSON));
    }

    /**
     * list carrier entities
     * List carrier entities
     */
    async listCarrier(requestParameters: ListCarrierRequest): Promise<Array<EntCarrier>> {
        const response = await this.listCarrierRaw(requestParameters);
        return await response.value();
    }

    /**
     * list contagious entities
     * List contagious entities
     */
    async listContagiousRaw(requestParameters: ListContagiousRequest): Promise<runtime.ApiResponse<Array<EntContagious>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/contagiouss`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntContagiousFromJSON));
    }

    /**
     * list contagious entities
     * List contagious entities
     */
    async listContagious(requestParameters: ListContagiousRequest): Promise<Array<EntContagious>> {
        const response = await this.listContagiousRaw(requestParameters);
        return await response.value();
    }

    /**
     * list employee entities
     * List employee entities
     */
    async listEmployeeRaw(requestParameters: ListEmployeeRequest): Promise<runtime.ApiResponse<Array<EntEmployee>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntEmployeeFromJSON));
    }

    /**
     * list employee entities
     * List employee entities
     */
    async listEmployee(requestParameters: ListEmployeeRequest): Promise<Array<EntEmployee>> {
        const response = await this.listEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * list patient entities
     * List patient entities
     */
    async listPatientRaw(requestParameters: ListPatientRequest): Promise<runtime.ApiResponse<Array<EntPatient>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntPatientFromJSON));
    }

    /**
     * list patient entities
     * List patient entities
     */
    async listPatient(requestParameters: ListPatientRequest): Promise<Array<EntPatient>> {
        const response = await this.listPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * list statistic entities
     * List statistic entities
     */
    async listStatisticRaw(requestParameters: ListStatisticRequest): Promise<runtime.ApiResponse<Array<EntStatistic>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/statistics`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntStatisticFromJSON));
    }

    /**
     * list statistic entities
     * List statistic entities
     */
    async listStatistic(requestParameters: ListStatisticRequest): Promise<Array<EntStatistic>> {
        const response = await this.listStatisticRaw(requestParameters);
        return await response.value();
    }

}
