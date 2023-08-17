.. procedure::
   :style: normal

   .. step:: Start the daemon.

      Run the following command to start the daemon:

      .. code-block::

         docker run -d --name mongodb/atlas mongodb/atlas

   .. step:: Get a shell.

      Run the following command to get a shell with an environment file:
      
      .. code-block:: 
          
         docker exec --env-file atlas.env --rm -it mongodb/atlas bash

   .. step:: Authenticate and run {+atlas-cli+} commands.

      To authenticate and run commands, follow the steps to :ref:`connect from the {+atlas-cli+}
      <connect-atlas-cli>`, adding 
      ``docker exec --env-file ./atlas.env --rm mongodb/atlas`` before
      each {+atlas-cli+} command. For example, to run the 
      :ref:`atlas config init <atlas-config-init>` command with an
      environment variable file, run the following command, replacing
      ``atlas.env`` with the name of the environment file:
      
      .. code-block::
         
         docker exec --env-file ./atlas.env --rm mongodb/atlas atlas config init

      .. note::

         The ``atlas config init`` command requires setting API keys in the `environment variable file 
         <https://docs.docker.com/engine/reference/commandline/run/#env>`__. 
         To learn more, see 
         :ref:`{+atlas-cli+} environment variables
         <atlas-cli-env-vars>`.

      After you authenticate, you can run {+atlas-cli+} commands by
      adding adding 
      ``docker exec --env-file ./atlas.env --rm mongodb/atlas`` before
      each {+atlas-cli+} command. For example, to run the 
      :ref:`atlas --help <atlas>` command with an environment variable
      file, run the following command, replacing ``atlas.env`` with the
      name of the environment file:

      .. code-block::

         docker exec --env-file ./atlas.env --rm mongodb/atlas atlas --help
