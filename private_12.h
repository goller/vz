//
//  private_12.h
//
//  Created by Christophe Fergeau
//

#pragma once

#import "virtualization_helper.h"
#import <Virtualization/Virtualization.h>

/* Virtualization Framework Private API */
//NS_ASSUME_NONNULL_BEGIN

VZ_EXPORT API_AVAILABLE(macos(12.0))
@interface _VZ16550SerialPortConfiguration : VZSerialPortConfiguration

- (instancetype)init NS_DESIGNATED_INITIALIZER;

@end

//NS_ASSUME_NONNULL_END

/* exported from cgo */

/* Configurations */
void *newVZ16550SerialPortConfiguration(void *attachment);
