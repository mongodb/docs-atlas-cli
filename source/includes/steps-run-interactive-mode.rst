.. procedure::
   :style: normal

   .. step:: Get a shell.

      Run the command to get a shell in interactive mode.
      
      Without an Environment Variable File
      ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

      To get a shell without an environment variable file, run the
      following command:
    
      .. code-block:: 
          
         docker run --rm -it mongodb/atlas bash

      With an Environment Variable File
      ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

      To get a shell using an environment file, run the following
      command, replacing ``atlas.env`` with the name of the environment
      file:

      .. code-block::

         docker run --env-file atlas.env --rm -it mongodb/atlas bash

   .. step:: Authenticate and run {+atlas-cli+} commands.

      Without an Environment Variable File
      ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

      To authenticate and run commands without an environment variable
      file, follow the steps to :ref:`connect from the {+atlas-cli+}
      <connect-atlas-cli>`. For example, you can run 
      :ref:`atlas auth login <atlas-auth-login>` to authenticate: 

      .. code-block::

         atlas auth login
      
      After you authenticate, you can run Atlas CLI commands. For
      example, you can run :ref:`atlas --help <atlas>` to learn about
      available commands:

      .. code-block::

         atlas --help

      With an Environment Variable File
      ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

      To authenticate and run commands with an environment variable
      file, follow the steps to :ref:`connect from the {+atlas-cli+}
      <connect-atlas-cli>`, adding 
      ``docker run --env-file ./atlas.env --rm mongodb/atlas`` before
      each {+atlas-cli+} command. For example, to run the 
      :ref:`atlas config init <atlas-config-init>` command with an
      environment variable file, run the following command, replacing
      ``atlas.env`` with the name of the environment file:
      
      .. code-block::
         
         docker run --env-file ./atlas.env --rm mongodb/atlas atlas config init

      .. note::

         The ``atlas config init`` command requires setting API keys in the `environment variable file 
         <https://docs.docker.com/engine/reference/commandline/run/#env>`__. 
         To learn more, see 
         :ref:`{+atlas-cli+} environment variables
         <atlas-cli-env-vars>`.

      After you authenticate, you can run {+atlas-cli+} commands by
      adding adding 
      ``docker run --env-file ./atlas.env --rm mongodb/atlas`` before
      each {+atlas-cli+} command. For example, to run the 
      :ref:`atlas --help <atlas>` command with an environment variable
      file, run the following command, replacing ``atlas.env`` with the
      name of the environment file:

      .. code-block::

         docker run --env-file ./atlas.env --rm mongodb/atlas atlas --help
