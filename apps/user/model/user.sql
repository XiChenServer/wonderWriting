-- 创建名为 Users 的表
CREATE TABLE Users (
    -- 用户ID，主键，自增长
                       UserID INT PRIMARY KEY AUTO_INCREMENT,
    -- 用户昵称，不能为空
                       Nickname VARCHAR(255) NOT NULL,
    -- 用户账号，唯一且不能为空
                       Username VARCHAR(50) UNIQUE NOT NULL,
    -- 用户邮箱，唯一，可以为空，默认值为NULL
                       Email VARCHAR(255) UNIQUE DEFAULT NULL,
    -- 用户手机号，唯一且不能为空
                       Phone VARCHAR(20) UNIQUE NOT NULL,
    -- 用户密码，不能为空
                       Password VARCHAR(255) NOT NULL,
    -- 注册时间，默认为当前时间
                       RegistrationTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- 最后登录时间，更新时自动更新为当前时间
                       LastLoginTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 用户状态，枚举类型，默认为激活状态
                       Status ENUM('Active', 'Inactive') DEFAULT 'Active',
    -- 用户角色，枚举类型，默认为普通用户角色
                       Role ENUM('User', 'Admin') DEFAULT 'User',
                       PRIMARY KEY (`UserID`)
);




