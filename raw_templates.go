package main

var rawTemplates = map[string]string{"templates/bin/wp.tmpl": "#!/bin/sh\nset -e\n\nFLAGS=\n\n# Add -t flag iff STDIN is a TTY\nif [ -t 0 ]; then\n  FLAGS=-t\nfi\n\nCONTAINER=`docker-compose ps -q wordpress`\n\n# We can't use docker-compose here because docker-compose exec is equivalent\n# to docker exec -ti and docker-compose exec -T is equivalent to\n# docker exec. There is no docker-compose equivalent to docker exec -i.\n#\n# Issue: https://github.com/docker/compose/issues/3352\n\ndocker exec -i ${FLAGS} ${CONTAINER} wp \"${@}\"\n", "templates/config/server.php.tmpl": "<?php\nif(!defined('MULTISITE')) {\n    define( 'MULTISITE', true );\n}\nif(!defined('SUBDOMAIN_INSTALL')) {\n    define( 'SUBDOMAIN_INSTALL', false );\n}\nif(!defined('DOMAIN_CURRENT_SITE')) {\n    define( 'DOMAIN_CURRENT_SITE', 'localhost' );\n}\nif(!defined('PATH_CURRENT_SITE')) {\n    define( 'PATH_CURRENT_SITE', '/' );\n}\nif(!defined('SITE_ID_CURRENT_SITE')) {\n    define( 'SITE_ID_CURRENT_SITE', 1 );\n}\nif(!defined('BLOG_ID_CURRENT_SITE')) {\n    define( 'BLOG_ID_CURRENT_SITE', 1 );\n}\n", "templates/docker-compose.yml.tmpl": "version: \"3\"\n\nvolumes:\n  mysql_data_{{.Name}}:\n\nservices:\n  mailcatcher:\n    image: schickling/mailcatcher\n    ports:\n      - \"1080:1080\"\n\n  beanstalk:\n    image: schickling/beanstalkd\n    ports:\n      - \"11300:11300\"\n\n  beanstalkd_console:\n    image: agaveapi/beanstalkd-console\n    ports:\n      - \"2080:80\"\n    environment:\n      BEANSTALKD_HOST: beanstalk\n      BEANSTALKD_PORT: 11300\n\n  mysql:\n    image: mariadb:10\n    ports:\n      - \"3306:3306\"\n    volumes:\n      - mysql_data_{{.Name}}:/var/lib/mysql\n    environment:\n      MYSQL_DATABASE: wordpress\n      MYSQL_ROOT_PASSWORD: foobar\n\n  wordpress:\n    image: thedxw/wpc-wordpress\n    ports:\n      - \"80:80\"\n    links:\n      - mysql\n      - mailcatcher\n      - beanstalk\n    volumes:\n      - .:/usr/src/app\n      - ./wp-content:/var/www/html/wp-content\n", "templates/setup/external.sh.tmpl": "#!/bin/sh\nset -e\n\n##\n## This code will be run during setup, OUTSIDE the container.\n##\n## Because `whippet` running inside the container wouldn't be able to connect\n## to private repositories.\n##\n\nif test -f whippet.json; then\n  whippet deps install\nfi\n", "templates/setup/internal.sh.tmpl": "#!/bin/sh\nset -e\n\n##\n## This code will be run during setup, INSIDE the container.\n##\n\n##############\n#\u00a0Config\n##############\ntitle=\"Your site title here\"\ntheme=your-theme-slug\nplugins=\"a-space-separated list-of plugins-to-activate\"\ncontent=/usr/src/app/setup/content\n\nwp core {{.InstallType}} --skip-email --admin_user=admin --admin_password=admin --admin_email=admin@localhost.invalid --url=http://localhost --title=\"$title\"\n\nfor plugin in $plugins\ndo\n  if wp plugin is-installed $plugin\n  then\n    wp plugin activate $plugin {{.ActivationType}}\n  else\n      echo \"\\033[96mWarning:\\033[0m Plugin '\"$plugin\"' could not be found. Have you installed it?\"\n  fi\ndone\n\nif wp theme is-installed $theme\nthen\n  {{.ThemeEnable}}\n  wp theme activate $theme\nelse\n  echo \"\\033[96mWarning:\\033[0m Theme '\"$theme\"' could not be found. Have you installed it?\"\nfi\n\nimport() {\n  for file in $content/*.xml\n  do\n    echo \"Importing $file...\"\n    wp import $file --authors=skip\n  done\n}\n\nif [ \"$(ls -A $content)\" ]\nthen\n  if wp plugin is-installed wordpress-importer\n  then\n    wp plugin activate wordpress-importer\n    import\n  else\n    echo \"WordPress Importer not installed... installing now\"\n    wp plugin install wordpress-importer --activate\n    import\n    wp plugin uninstall wordpress-importer --deactivate\n  fi\nelse\n  echo \"No content to be imported\"\nfi\n", "templates/bin/console.tmpl": "#!/bin/sh\nset -e\n\nexec docker-compose exec wordpress bash\n", "templates/bin/setup.tmpl": "#!/bin/sh\nset -e\n#\n# Runs all site setup scripts\n\n`dirname $0`/../setup/external.sh\ndocker-compose exec wordpress /usr/src/app/setup/internal.sh\n"}