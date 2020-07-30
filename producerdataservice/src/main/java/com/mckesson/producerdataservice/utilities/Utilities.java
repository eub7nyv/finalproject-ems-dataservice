package com.mckesson.producerdataservice.utilities;

import org.apache.commons.lang3.StringUtils;

/**
 * @author anoopunnikrishnan
 *
 */

public class Utilities {
    public static String environmentOrDefault(String envKey, String defaultVal) {
        String val = StringUtils.trimToNull(System.getenv(envKey));
        return (null == val) ? defaultVal : val;
    }
}