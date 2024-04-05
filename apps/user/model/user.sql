-- 创建名为 Users 的表
CREATE TABLE Users (
                       UserID INT PRIMARY KEY AUTO_INCREMENT, -- 用户ID，主键，自增长
                       Nickname VARCHAR(255) NOT NULL, -- 用户昵称，不能为空
                       Account VARCHAR(50) UNIQUE NOT NULL, -- 用户账号，唯一且不能为空
                       Email VARCHAR(255) UNIQUE DEFAULT NULL, -- 用户邮箱，唯一，可以为空，默认值为NULL
                       Phone VARCHAR(20) UNIQUE NOT NULL, -- 用户手机号，唯一且不能为空
                       Password VARCHAR(255) NOT NULL, -- 用户密码，不能为空
                       RegistrationTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 注册时间，默认为当前时间
                       LastLoginTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 最后登录时间，更新时自动更新为当前时间
                       Status ENUM('Active', 'Inactive') DEFAULT 'Active', -- 用户状态，枚举类型，默认为激活状态
                       Role ENUM('User', 'Admin') DEFAULT 'User', -- 用户角色，枚举类型，默认为普通用户角色
                       BackgroundImage VARCHAR(255) DEFAULT NULL, -- 用户背景图片，可以为空，默认为NULL
                       AvatarBackground VARCHAR(255) DEFAULT NULL, -- 头像背景图片，可以为空，默认为NULL
                       PostCount INT DEFAULT 0, -- 帖子数，默认为0
                       FollowCount INT DEFAULT 0, -- 关注数，默认为0
                       FansCount INT DEFAULT 0, -- 粉丝数，默认为0
                       LikeCount INT DEFAULT 0, -- 获赞数，默认为0
                       PointCount INT DEFAULT 0, -- 积分数，默认为0
                       PRIMARY KEY (`UserID`)
);




