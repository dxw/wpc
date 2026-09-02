<?php

// Per-project configuration
if (file_exists('/usr/src/app/config/server.php')) {
    require('/usr/src/app/config/server.php');
}

// Per-host configuration (for setting WP_SITEURL to http://x.local etc)
if (file_exists('/usr/src/app/config/server-local.php')) {
    require('/usr/src/app/config/server-local.php');
}

// For some reason this is not being set correctly by default
if (!defined('DB_CHARSET')) {
    define('DB_CHARSET', $_ENV['DB_CHARSET'] ?? 'utf8mb4');
}

// Set URLs if they aren't set already
$host = isset($_SERVER['SERVER_NAME']) ? $_SERVER['SERVER_NAME'] : 'localhost';
if (!defined('WP_SITEURL')) {
    define('WP_SITEURL', $_ENV['WP_SITEURL'] ?? 'http://'.$host);
}
if (!defined('WP_HOME')) {
    define('WP_HOME', $_ENV['WP_HOME'] ?? 'http://'.$host);
}

// beanstalk
define('BEANSTALKD_HOST', $_ENV['BEANSTALKD_HOST'] ?? 'beanstalk');

// mu-plugins
define('WPMU_PLUGIN_DIR', $_ENV['WPMU_PLUGIN_DIR'] ?? '/usr/src/mu-plugins');

// Database
define('DB_HOST', $_ENV['DB_HOST'] ?? 'mysql:3306');
define('DB_NAME', $_ENV['DB_NAME'] ?? 'wordpress');
define('DB_USER', $_ENV['DB_USER'] ?? 'root');
define('DB_PASSWORD', $_ENV['DB_PASSWORD'] ?? 'foobar');

// wp-config stuff

if (!defined('WP_DEBUG')) {
    define('WP_DEBUG', $_ENV['WP_DEBUG'] ?? false);
}
define('WP_DEBUG_DISPLAY', $_ENV['WP_DEBUG_DISPLAY'] ?? true);
define('WP_ALLOW_MULTISITE', $_ENV['WP_ALLOW_MULTISITE'] ?? true);
define('FS_METHOD', $_ENV['FS_METHOD'] ?? 'direct');
define('CORE_UPGRADE_SKIP_NEW_BUNDLED', $_ENV['CORE_UPGRADE_SKIP_NEW_BUNDLED'] ?? true);

define('AUTH_KEY',         $_ENV['AUTH_KEY'] ?? 'put your unique phrase here');
define('SECURE_AUTH_KEY',  $_ENV['SECURE_AUTH_KEY'] ?? 'put your unique phrase here');
define('LOGGED_IN_KEY',    $_ENV['LOGGED_IN_KEY'] ?? 'put your unique phrase here');
define('NONCE_KEY',        $_ENV['NONCE_KEY'] ?? 'put your unique phrase here');
define('AUTH_SALT',        $_ENV['AUTH_SALT'] ?? 'put your unique phrase here');
define('SECURE_AUTH_SALT', $_ENV['SECURE_AUTH_SALT'] ?? 'put your unique phrase here');
define('LOGGED_IN_SALT',   $_ENV['LOGGED_IN_SALT'] ?? 'put your unique phrase here');
define('NONCE_SALT',       $_ENV['NONCE_SALT'] ?? 'put your unique phrase here');

define('WPLANG', $_ENV['WPLANG'] ?? '');
if (!defined('ABSPATH')) {
    define('ABSPATH', $_ENV['ABSPATH'] ?? dirname(__FILE__) . '/');
}
$table_prefix = $_ENV['TABLE_PREFIX'] ?? 'wp_';

require_once(ABSPATH . 'wp-settings.php');
