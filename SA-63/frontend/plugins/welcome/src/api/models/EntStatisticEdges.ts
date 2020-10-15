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

import { exists, mapValues } from '../runtime';
import {
    EntArea,
    EntAreaFromJSON,
    EntAreaFromJSONTyped,
    EntAreaToJSON,
    EntCarrier,
    EntCarrierFromJSON,
    EntCarrierFromJSONTyped,
    EntCarrierToJSON,
    EntContagious,
    EntContagiousFromJSON,
    EntContagiousFromJSONTyped,
    EntContagiousToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatisticEdges
 */
export interface EntStatisticEdges {
    /**
     * 
     * @type {EntArea}
     * @memberof EntStatisticEdges
     */
    area?: EntArea;
    /**
     * 
     * @type {EntCarrier}
     * @memberof EntStatisticEdges
     */
    carrier?: EntCarrier;
    /**
     * 
     * @type {EntContagious}
     * @memberof EntStatisticEdges
     */
    contagious?: EntContagious;
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntStatisticEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntStatisticEdges
     */
    patient?: EntPatient;
}

export function EntStatisticEdgesFromJSON(json: any): EntStatisticEdges {
    return EntStatisticEdgesFromJSONTyped(json, false);
}

export function EntStatisticEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatisticEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'area': !exists(json, 'area') ? undefined : EntAreaFromJSON(json['area']),
        'carrier': !exists(json, 'carrier') ? undefined : EntCarrierFromJSON(json['carrier']),
        'contagious': !exists(json, 'contagious') ? undefined : EntContagiousFromJSON(json['contagious']),
        'employee': !exists(json, 'employee') ? undefined : EntEmployeeFromJSON(json['employee']),
        'patient': !exists(json, 'patient') ? undefined : EntPatientFromJSON(json['patient']),
    };
}

export function EntStatisticEdgesToJSON(value?: EntStatisticEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'area': EntAreaToJSON(value.area),
        'carrier': EntCarrierToJSON(value.carrier),
        'contagious': EntContagiousToJSON(value.contagious),
        'employee': EntEmployeeToJSON(value.employee),
        'patient': EntPatientToJSON(value.patient),
    };
}

