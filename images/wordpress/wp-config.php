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
    define('DB_CHARSET', 'utf8mb4');
}

// Set URLs if they aren't set already
$host = isset($_SERVER['SERVER_NAME']) ? $_SERVER['SERVER_NAME'] : 'localhost';
if (!defined('WP_SITEURL')) {
    define('WP_SITEURL', 'http://'.$host);
}
if (!defined('WP_HOME')) {
    define('WP_HOME', 'http://'.$host);
}

// beanstalk
define('BEANSTALKD_HOST', 'beanstalk');

// mu-plugins
define('WPMU_PLUGIN_DIR', '/usr/src/mu-plugins');

// Database
define('DB_HOST', 'mysql:3306');
define('DB_NAME', 'wordpress');
define('DB_USER', 'root');
define('DB_PASSWORD', 'foobar');

// wp-config stuff

if (!defined('WP_DEBUG')) {
    define('WP_DEBUG', false);
}
define('WP_DEBUG_DISPLAY', true);
define('WP_ALLOW_MULTISITE', true);
define('FS_METHOD', 'direct');

define('AUTH_KEY',         'put your unique phrase here');
define('SECURE_AUTH_KEY',  'put your unique phrase here');
define('LOGGED_IN_KEY',    'put your unique phrase here');
define('NONCE_KEY',        'put your unique phrase here');
define('AUTH_SALT',        'put your unique phrase here');
define('SECURE_AUTH_SALT', 'put your unique phrase here');
define('LOGGED_IN_SALT',   'put your unique phrase here');
define('NONCE_SALT',       'put your unique phrase here');

define('WPLANG', '');
if (!defined('ABSPATH')) {
    define('ABSPATH', dirname(__FILE__) . '/');
}
$table_prefix = 'wp_';

require_once(ABSPATH . 'wp-settings.php');
